import dagre from 'dagre';
import { Node, Edge, MarkerType, Position } from 'reactflow';

// Tipagem baseada no seu JSON
type ServiceData = {
  name: string;
  system: string;
  team: string;
  domain: string;
  dependencies: string[];
};

const nodeWidth = 260;
const nodeHeight = 120;

export const getLayoutedElements = (
  allServices: ServiceData[], 
  selectedTeam: string, 
  maxDepth: number
) => {
  const dagreGraph = new dagre.graphlib.Graph();
  dagreGraph.setDefaultEdgeLabel(() => ({}));

  // Configuração da direção do grafo (Left to Right parece ser o seu desenho)
  dagreGraph.setGraph({ rankdir: 'LR' });

  const nodes: Node[] = [];
  const edges: Edge[] = [];
  const visited = new Set<string>();

  // 1. Encontrar nós raiz (do time selecionado)
  const rootServices = allServices.filter(s => s.team === selectedTeam);
  
  // Função recursiva para buscar dependências até o nível X
  const buildGraph = (serviceName: string, currentDepth: number) => {
    if (visited.has(serviceName)) return;
    visited.add(serviceName);

    const service = allServices.find(s => s.name === serviceName);
    if (!service) return;

    const isFocusTeam = service.team === selectedTeam;

    // Adiciona o Nó
    nodes.push({
      id: service.name,
      type: 'serviceNode', // Nome do nosso componente customizado
      data: { 
        label: service.name,
        system: service.system,
        team: service.team,
        domain: service.domain,
        isFocusTeam: isFocusTeam
      },
      position: { x: 0, y: 0 }, // Dagre vai calcular isso depois
    });

    // Se atingiu o limite de profundidade e NÃO é um nó raiz, paramos aqui
    // (A lógica de profundidade geralmente se aplica a partir dos nós selecionados)
    if (!isFocusTeam && currentDepth >= maxDepth) return;

    // Processar dependências
    service.dependencies.forEach(depName => {
      // Adiciona a conexão (Edge)
      edges.push({
        id: `${service.name}-${depName}`,
        source: service.name,
        target: depName,
        type: 'smoothstep',
        markerEnd: { type: MarkerType.ArrowClosed },
        style: { stroke: isFocusTeam ? '#3b82f6' : '#9ca3af' }
      });

      // Recurso: incrementa profundidade se sairmos do time selecionado
      const nextDepth = isFocusTeam ? 0 : currentDepth + 1;
      buildGraph(depName, nextDepth);
    });
  };

  // Inicia a construção para cada serviço do time selecionado
  rootServices.forEach(s => buildGraph(s.name, 0));

  // 2. Calcular Layout com Dagre
  nodes.forEach((node) => {
    dagreGraph.setNode(node.id, { width: nodeWidth, height: nodeHeight });
  });

  edges.forEach((edge) => {
    dagreGraph.setEdge(edge.source, edge.target);
  });

  dagre.layout(dagreGraph);

  // 3. Aplicar posições calculadas aos nós
  const layoutedNodes = nodes.map((node) => {
    const nodeWithPosition = dagreGraph.node(node.id);
    node.targetPosition = Position.Left;
    node.sourcePosition = Position.Right;
    // Precisamos ajustar o ponto central, pois o dagre considera o centro, e o react flow o topo-esquerdo
    node.position = {
      x: nodeWithPosition.x - nodeWidth / 2,
      y: nodeWithPosition.y - nodeHeight / 2,
    };
    return node;
  });

  return { nodes: layoutedNodes, edges };
};