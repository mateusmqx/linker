import React, { memo } from 'react';
import { Handle, Position, NodeProps } from 'reactflow';

const ServiceNode = ({ data }: NodeProps) => {
  // Define a cor baseada se o nó é o alvo principal da busca ou apenas uma dependência
  const isHighlight = data.isHighlight;
  const borderColor = isHighlight ? 'border-orange-500' : 'border-gray-500';
  const textColor = isHighlight ? 'text-orange-500' : 'text-gray-200';
  const labelColor = isHighlight ? 'text-orange-400' : 'text-gray-400';

  return (
    <div className={`w-[250px] bg-[#141414] rounded-lg border-2 ${borderColor} shadow-lg transition-all duration-300`}>
      {/* Handle de Entrada (Esquerda) */}
      <Handle type="target" position={Position.Left} className="!bg-gray-500 !w-3 !h-3" />
      
      <div className="p-3 flex flex-col h-full min-h-[100px] justify-between">
        
        {/* Header: Nome */}
        <div className="mb-2">
          <span className={`text-xs font-mono uppercase ${labelColor} opacity-70`}>Name</span>
          <div className={`font-bold text-sm break-words ${textColor}`}>
            {data.label}
          </div>
        </div>

        {/* Body: Time */}
        <div className="mb-4">
          <span className="text-[10px] text-gray-500 block uppercase">Team</span>
          <span className="text-xs text-gray-300 bg-gray-800 px-2 py-1 rounded">
            {data.team}
          </span>
        </div>

        {/* Footer: Links */}
        <div className="flex justify-between border-t border-gray-800 pt-2 mt-auto">
          <span className="text-[10px] text-gray-600 cursor-pointer hover:text-white transition-colors">
            repository
          </span>
          <span className="text-[10px] text-gray-600 cursor-pointer hover:text-white transition-colors">
            documentation
          </span>
        </div>
      </div>

      {/* Handle de Saída (Direita) */}
      <Handle type="source" position={Position.Right} className="!bg-gray-500 !w-3 !h-3" />
    </div>
  );
};

export default memo(ServiceNode);