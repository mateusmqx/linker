"use client";

import React, { useMemo, useCallback, useState } from 'react';
import ReactFlow, {
  Background,
  Controls,
  MiniMap,
  useNodesState,
  useEdgesState,
  Node,
  Edge,
  ConnectionLineType,
  Panel,
} from 'reactflow';
import 'reactflow/dist/style.css'; // Importante!
import { getLayoutedElements } from '@/utils/layout';
import rawData from '@/data/services.json'; // Importe seu JSON

// Tipagem dos dados originais
type ServiceData = {
  name: string;
  system: string;
  team: string;
  domain: string;
  dependencies: string[];
};

export default function ServiceFlow() {
  // 1. Transformar o JSON em Nodes e Edges iniciais
  const { nodes: initialNodes, edges: initialEdges } = useMemo(() => {
    const nodes: Node[] = [];
    const edges: Edge[] = [];
    
    // Mapeia cores por domínio para ficar bonito
    const domainColors: Record<string, string> = {
      risk: '#ef4444', // red
      'core-banking': '#a855f7', // purple
      payments: '#3b82f6', // blue
      platform: '#64748b', // slate
      // Adicione outros conforme necessário
    };

    (rawData as ServiceData[]).forEach((service) => {
      // Cria o Nó
      nodes.push({
        id: service.name,
        data: { label: service.name, details: service }, // Passamos o objeto todo para usar no clique
        position: { x: 0, y: 0 }, // Será calculado pelo layout depois
        style: {
          background: '#1e293b',
          color: '#fff',
          border: '1px solid #334155',
          borderRadius: '8px',
          width: 250,
          padding: '10px',
          borderLeft: `5px solid ${domainColors[service.domain] || '#94a3b8'}`,
        },
      });

      // Cria as conexões (Arestas)
      service.dependencies.forEach((dep) => {
        edges.push({
          id: `${service.name}-${dep}`,
          source: service.name,
          target: dep,
          type: 'smoothstep', // Linhas retas e curvas suaves (estilo circuito)
          animated: true,
          style: { stroke: '#475569' },
        });
      });
    });

    return getLayoutedElements(nodes, edges);
  }, []);

  const [nodes, setNodes, onNodesChange] = useNodesState(initialNodes);
  const [edges, setEdges, onEdgesChange] = useEdgesState(initialEdges);
  const [selectedService, setSelectedService] = useState<ServiceData | null>(null);

  // Função executada ao clicar no nó
  const onNodeClick = useCallback((event: React.MouseEvent, node: Node) => {
    setSelectedService(node.data.details);
  }, []);

  return (
    <div className="flex h-screen w-full bg-slate-950">
      {/* Área do Grafo */}
      <div className="flex-grow h-full relative">
        <ReactFlow
          nodes={nodes}
          edges={edges}
          onNodesChange={onNodesChange}
          onEdgesChange={onEdgesChange}
          onNodeClick={onNodeClick}
          fitView
          minZoom={0.1}
        >
          <Background color="#334155" gap={16} />
          <Controls className="bg-white" />
          <MiniMap 
            nodeColor={() => '#475569'} 
            maskColor="rgba(0,0,0, 0.6)"
            style={{ backgroundColor: '#1e293b' }}
          />
          
          <Panel position="top-left" className="bg-slate-800 p-2 rounded text-white border border-slate-700">
            Total de Microserviços: <strong>{nodes.length}</strong>
          </Panel>
        </ReactFlow>
      </div>

      {/* Sidebar de Detalhes (Aparece ao clicar) */}
      {selectedService && (
        <div className="w-96 bg-slate-900 border-l border-slate-700 p-6 shadow-2xl overflow-y-auto z-10 animate-slide-in">
          <div className="flex justify-between items-start mb-6">
            <h2 className="text-xl font-bold text-blue-400 break-words w-full">
              {selectedService.name}
            </h2>
            <button 
              onClick={() => setSelectedService(null)}
              className="text-slate-400 hover:text-white"
            >
              ✕
            </button>
          </div>

          <div className="space-y-6">
            <div>
              <label className="text-xs uppercase text-slate-500 font-semibold">Sistema</label>
              <p className="text-slate-200">{selectedService.system}</p>
            </div>
            
            <div className="grid grid-cols-2 gap-4">
              <div>
                <label className="text-xs uppercase text-slate-500 font-semibold">Time</label>
                <span className="inline-block px-2 py-1 mt-1 rounded bg-blue-900/50 text-blue-300 text-sm border border-blue-800">
                  {selectedService.team}
                </span>
              </div>
              <div>
                <label className="text-xs uppercase text-slate-500 font-semibold">Domínio</label>
                <span className="inline-block px-2 py-1 mt-1 rounded bg-purple-900/50 text-purple-300 text-sm border border-purple-800">
                  {selectedService.domain}
                </span>
              </div>
            </div>

            <div>
              <label className="text-xs uppercase text-slate-500 font-semibold">Dependências ({selectedService.dependencies.length})</label>
              {selectedService.dependencies.length > 0 ? (
                <ul className="mt-2 space-y-1">
                  {selectedService.dependencies.map(dep => (
                    <li key={dep} className="text-sm text-slate-300 flex items-center gap-2">
                      <span className="w-1.5 h-1.5 rounded-full bg-slate-500"></span>
                      {dep}
                    </li>
                  ))}
                </ul>
              ) : (
                <p className="text-sm text-slate-500 italic mt-1">Nenhuma dependência direta.</p>
              )}
            </div>

            <div className="pt-6 border-t border-slate-800 flex flex-col gap-2">
               {/* Exemplo de botões de ação */}
               <button className="w-full py-2 bg-slate-800 hover:bg-slate-700 text-white rounded text-sm transition">
                 Ver no GitHub
               </button>
               <button className="w-full py-2 bg-slate-800 hover:bg-slate-700 text-white rounded text-sm transition">
                 Ver Documentação
               </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}