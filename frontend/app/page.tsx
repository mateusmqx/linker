"use client"

import React, { useState, useCallback, useEffect } from 'react';
import ReactFlow, { 
  Background, 
  Controls, 
  useNodesState, 
  useEdgesState 
} from 'reactflow';
import 'reactflow/dist/style.css';

import ServiceNode from './ServiceNode';
import { getLayoutedElements } from './GraphBuilder';
import rawData from '../data/services.json'; // O seu JSON gigante

const nodeTypes = { serviceNode: ServiceNode };

export default function ArchitectureMap() {
  const [nodes, setNodes, onNodesChange] = useNodesState([]);
  const [edges, setEdges, onEdgesChange] = useEdgesState([]);
  
  // Estados para os filtros
  const [selectedTeam, setSelectedTeam] = useState('fraud-prevention');
  const [depthLevel, setDepthLevel] = useState(1);

  // Extrair lista única de times para o select
  const teams = Array.from(new Set(rawData.map(s => s.team))).sort();

  // Função para atualizar o gráfico
  const refreshGraph = useCallback(() => {
    const { nodes: layoutedNodes, edges: layoutedEdges } = getLayoutedElements(
      rawData, 
      selectedTeam, 
      depthLevel
    );
    setNodes(layoutedNodes);
    setEdges(layoutedEdges);
  }, [selectedTeam, depthLevel, setNodes, setEdges]);

  // Atualiza sempre que mudar o filtro
  useEffect(() => {
    refreshGraph();
  }, [refreshGraph]);

  return (
    <div style={{ width: '100vw', height: '100vh', display: 'flex', flexDirection: 'column' }}>
      
      {/* Barra de Filtros (Simulando seu desenho) */}
      <div style={{ padding: '20px', borderBottom: '1px solid #ccc', background: '#f9f9f9', display: 'flex', gap: '20px', alignItems: 'center' }}>
        <div>
          <label style={{ display: 'block', fontWeight: 'bold', fontSize: '12px' }}>Team</label>
          <select 
            value={selectedTeam} 
            onChange={(e) => setSelectedTeam(e.target.value)}
            style={{ padding: '5px', borderRadius: '4px', border: '1px solid #ccc' }}
          >
            {teams.map(t => <option key={t} value={t}>{t}</option>)}
          </select>
        </div>

        <div>
          <label style={{ display: 'block', fontWeight: 'bold', fontSize: '12px' }}>External Deep Level</label>
          <input 
            type="number" 
            min="0" 
            max="5" 
            value={depthLevel} 
            onChange={(e) => setDepthLevel(Number(e.target.value))}
            style={{ padding: '5px', width: '60px', borderRadius: '4px', border: '1px solid #ccc' }}
          />
        </div>
      </div>

      {/* Área do React Flow */}
      <div style={{ flex: 1 }}>
        <ReactFlow
          nodes={nodes}
          edges={edges}
          onNodesChange={onNodesChange}
          onEdgesChange={onEdgesChange}
          nodeTypes={nodeTypes}
          fitView
        >
          <Background color="#aaa" gap={16} />
          <Controls />
        </ReactFlow>
      </div>
    </div>
  );
}