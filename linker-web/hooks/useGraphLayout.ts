import { useMemo } from 'react';
import dagre from 'dagre';
import { Edge, Node, Position } from 'reactflow';

type ServiceData = {
  name: string;
  system: string;
  team: string;
  domain: string;
  dependencies: string[];
};

const NODE_WIDTH = 250;
const NODE_HEIGHT = 150; // Ajustado para bater com o tamanho real do Card

export const useGraphLayout = (
  rawData: ServiceData[],
  filters: { team: string; searchTerm: string; depthIn: number; depthOut: number }
) => {
  const { nodes, edges } = useMemo(() => {
    if (!rawData || rawData.length === 0) return { nodes: [], edges: [] };

    // 1. Indexar dados
    const dataMap = new Map(rawData.map((item) => [item.name, item]));
    const parentMap: Record<string, string[]> = {};

    rawData.forEach((item) => {
      item.dependencies.forEach((dep) => {
        if (!parentMap[dep]) parentMap[dep] = [];
        parentMap[dep].push(item.name);
      });
    });

    // 2. Definir Raízes
    let rootIds: string[] = [];
    if (filters.searchTerm) {
      rootIds = rawData
        .filter((d) => d.name.toLowerCase().includes(filters.searchTerm.toLowerCase()))
        .map((d) => d.name);
    } else if (filters.team) {
      rootIds = rawData
        .filter((d) => d.team === filters.team)
        .map((d) => d.name);
    } else {
      return { nodes: [], edges: [] };
    }

    // 3. Travessia (BFS)
    const nodesToRender = new Set<string>();

    const expand = (currentIds: string[], depth: number, direction: 'IN' | 'OUT') => {
      let currentLayer = [...currentIds];
      let levels = depth;

      currentLayer.forEach((id) => nodesToRender.add(id));

      while (levels > 0 && currentLayer.length > 0) {
        const nextLayer: string[] = [];
        for (const id of currentLayer) {
          const neighbors =
            direction === 'OUT'
              ? dataMap.get(id)?.dependencies || []
              : parentMap[id] || [];

          for (const neighbor of neighbors) {
            // Só adiciona se o nó realmente existir no JSON (evita nós fantasmas)
            if (dataMap.has(neighbor)) {
              nodesToRender.add(neighbor);
              nextLayer.push(neighbor);
            }
          }
        }
        currentLayer = nextLayer;
        levels--;
      }
    };

    expand(rootIds, filters.depthOut, 'OUT');
    expand(rootIds, filters.depthIn, 'IN');

    // 4. Criar Nós e Arestas do React Flow
    const flowNodes: Node[] = [];
    const flowEdges: Edge[] = [];
    const renderedSet = Array.from(nodesToRender);

    renderedSet.forEach((id) => {
      const item = dataMap.get(id)!;
      const isRoot = rootIds.includes(id);

      flowNodes.push({
        id: item.name,
        type: 'serviceNode',
        data: {
          label: item.name,
          team: item.team,
          isHighlight: isRoot,
        },
        position: { x: 0, y: 0 },
      });

      item.dependencies.forEach((dep) => {
        if (nodesToRender.has(dep)) {
          flowEdges.push({
            id: `${id}-${dep}`,
            source: id,
            target: dep,
            animated: true,
            style: { stroke: isRoot ? '#f97316' : '#555' },
          });
        }
      });
    });

    // 5. Aplicar Layout Dagre (A CORREÇÃO ESTÁ AQUI)
    const dagreGraph = new dagre.graphlib.Graph();
    dagreGraph.setGraph({ rankdir: 'LR' });

    // CORREÇÃO: Define um objeto vazio padrão para as arestas. 
    // Sem isso, dagre tenta setar 'points' em undefined e quebra.
    dagreGraph.setDefaultEdgeLabel(() => ({})); 

    flowNodes.forEach((node) => {
      dagreGraph.setNode(node.id, { width: NODE_WIDTH, height: NODE_HEIGHT });
    });

    flowEdges.forEach((edge) => {
      dagreGraph.setEdge(edge.source, edge.target);
    });

    dagre.layout(dagreGraph);

    const layoutedNodes = flowNodes.map((node) => {
      const nodeWithPosition = dagreGraph.node(node.id);
      return {
        ...node,
        targetPosition: Position.Left,
        sourcePosition: Position.Right,
        position: {
          x: nodeWithPosition.x - NODE_WIDTH / 2,
          y: nodeWithPosition.y - NODE_HEIGHT / 2,
        },
      };
    });

    return { nodes: layoutedNodes, edges: flowEdges };
  }, [rawData, filters]);

  return { nodes, edges };
};