import React, { CSSProperties, memo } from 'react';
import { Handle, Position } from 'reactflow';

// Estilos básicos para o container
const nodeStyle : CSSProperties = {
  padding: '10px',
  borderRadius: '8px',
  border: '1px solid #777',
  width: '250px',
  fontSize: '12px',
  textAlign: 'center',
  backgroundColor: '#fff',
  boxShadow: '0 4px 6px -1px rgb(0 0 0 / 0.1)',
};

const ServiceNode = ({ data }: any) => {
  // Define a cor da borda/texto baseado se é do time selecionado ou dependência
  const borderColor = data.isFocusTeam ? '#3b82f6' : '#9ca3af'; // Azul ou Cinza
  const textColor = data.isFocusTeam ? '#1d4ed8' : '#4b5563';

  return (
    <div style={{ ...nodeStyle, borderColor: borderColor, color: textColor }}>
      {/* Handles para conectar as setas */}
      <Handle type="target" position={Position.Left} style={{ background: '#555' }} />
      
      <div style={{ paddingBottom: '15px', fontWeight: 'bold', fontSize: '14px' }}>
        {data.label}
      </div>
      
      <div style={{ display: 'flex', flexDirection: 'column', gap: '2px', fontSize: '10px', color: '#666' }}>
        <span>{data.system}</span>
        <span>{data.team}</span>
        <span>{data.domain}</span>
      </div>

      {/* Footer com links simula o desenho */}
      <div style={{ 
        display: 'flex', 
        justifyContent: 'space-between', 
        marginTop: '15px', 
        fontSize: '9px',
        borderTop: `1px solid ${borderColor}`,
        paddingTop: '4px'
      }}>
        <span style={{ cursor: 'pointer' }}>repository</span>
        <span style={{ cursor: 'pointer' }}>documentation</span>
      </div>

      <Handle type="source" position={Position.Right} style={{ background: '#555' }} />
    </div>
  );
};

export default memo(ServiceNode);