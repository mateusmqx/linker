'use client';

import React, { useState } from 'react';
import ReactFlow, { Background, Controls, NodeTypes } from 'reactflow';
import 'reactflow/dist/style.css';
import ServiceNode from '@/components/ServiceNode';
import { useGraphLayout } from '@/hooks/useGraphLayout';

// Seus dados JSON (cole aqui o array gigante que você mandou)
import rawData from './data.json'; 

// Registro do tipo de nó customizado
const nodeTypes: NodeTypes = {
  serviceNode: ServiceNode,
};

export default function MicroservicesMap() {
  // Estados dos Filtros
  const [searchTerm, setSearchTerm] = useState('');
  const [selectedTeam, setSelectedTeam] = useState('');
  const [depthIn, setDepthIn] = useState(1);
  const [depthOut, setDepthOut] = useState(1);

  // Extrair lista única de times para o select
  const teams = Array.from(new Set(rawData.map((d) => d.team))).sort();

  // Hook customizado que calcula o grafo
  const { nodes, edges } = useGraphLayout(rawData, {
    searchTerm,
    team: selectedTeam,
    depthIn,
    depthOut
  });

  return (
    <div className="w-screen h-screen bg-[#0a0a0a] text-white flex flex-col">
      
      {/* --- Barra de Filtros (Estilo do seu desenho) --- */}
      <div className="p-6 border-b border-gray-800 bg-[#111] grid grid-cols-1 md:grid-cols-4 gap-6 items-end z-10 shadow-xl">
        
        {/* Filtro: Nome da Aplicação */}
        <div className="flex flex-col gap-2">
            <label className="text-gray-400 font-handwriting text-sm">nome aplicação</label>
            <div className="flex border border-gray-600 rounded bg-[#1a1a1a]">
                <input 
                    type="text" 
                    placeholder="Ex: fraud-engine..."
                    className="bg-transparent p-2 outline-none text-white w-full placeholder-gray-600"
                    value={searchTerm}
                    onChange={(e) => {
                        setSearchTerm(e.target.value);
                        setSelectedTeam(''); // Limpa time se buscar por nome
                    }}
                />
                <button className="bg-gray-800 px-4 text-sm font-bold border-l border-gray-600">Buscar</button>
            </div>
        </div>

        <div className="flex items-center justify-center pb-2 text-gray-500 font-bold">OU</div>

        {/* Filtro: Time */}
        <div className="flex flex-col gap-2">
            <label className="text-gray-400 font-handwriting text-sm">Time</label>
            <div className="relative">
                <select 
                    className="w-full bg-[#1a1a1a] border border-gray-600 text-white p-2 rounded appearance-none cursor-pointer"
                    value={selectedTeam}
                    onChange={(e) => {
                        setSelectedTeam(e.target.value);
                        setSearchTerm(''); // Limpa nome se filtrar por time
                    }}
                >
                    <option value="">Selecione um time...</option>
                    {teams.map(team => <option key={team} value={team}>{team}</option>)}
                </select>
                <div className="absolute right-3 top-3 pointer-events-none">▼</div>
            </div>
        </div>

        {/* Filtros: Profundidade (IN/OUT) */}
        <div className="flex gap-4">
            <div className="flex flex-col gap-2 w-24">
                <label className="text-gray-400 font-handwriting text-sm text-center">In (pais)</label>
                <input 
                    type="number" 
                    min="0" 
                    max="5" 
                    className="bg-[#1a1a1a] border border-gray-600 p-2 rounded text-center text-white"
                    value={depthIn}
                    onChange={(e) => setDepthIn(Number(e.target.value))}
                />
            </div>
            <div className="flex flex-col gap-2 w-24">
                <label className="text-gray-400 font-handwriting text-sm text-center">Out (filhos)</label>
                <input 
                    type="number" 
                    min="0" 
                    max="5" 
                    className="bg-[#1a1a1a] border border-gray-600 p-2 rounded text-center text-white"
                    value={depthOut}
                    onChange={(e) => setDepthOut(Number(e.target.value))}
                />
            </div>
        </div>
      </div>

      {/* --- Área do Diagrama --- */}
      <div className="flex-1 w-full h-full relative">
        {nodes.length === 0 ? (
            <div className="absolute inset-0 flex items-center justify-center text-gray-600 flex-col">
                <p className="text-xl">Selecione um time ou busque um serviço para começar.</p>
                <p className="text-sm mt-2">O diagrama é gerado dinamicamente para evitar poluição visual.</p>
            </div>
        ) : (
            <ReactFlow
                nodes={nodes}
                edges={edges}
                nodeTypes={nodeTypes}
                fitView
                className="bg-[#0a0a0a]"
                minZoom={0.1}
            >
                <Background color="#222" gap={20} />
                <Controls className="bg-white text-black" />
            </ReactFlow>
        )}
      </div>
    </div>
  );
}