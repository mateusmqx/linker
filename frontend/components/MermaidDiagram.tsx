"use client";

import React, { useEffect, useRef, useState } from "react";
import mermaid from "mermaid";
import { TransformWrapper, TransformComponent, ReactZoomPanPinchContentRef } from "react-zoom-pan-pinch";

interface MermaidDiagramProps {
  chart: string;
}

const MermaidDiagram = ({ chart }: MermaidDiagramProps) => {
  const containerRef = useRef<HTMLDivElement>(null);
  const [isMounted, setIsMounted] = useState(false);
  const transformComponentRef = useRef<ReactZoomPanPinchContentRef>(null);

  useEffect(() => {
    setIsMounted(true);
    
    // Inicialização do Mermaid
    mermaid.initialize({
      startOnLoad: false,
      theme: "dark",
      securityLevel: "loose",
      // Ajustes para diagramas grandes
      flowchart: {
        useMaxWidth: false, // Importante: permite que o gráfico cresça além do container
        htmlLabels: true,
      },
    });
  }, []);

  useEffect(() => {
    const renderChart = async () => {
      if (containerRef.current && isMounted) {
        containerRef.current.innerHTML = "";
        const id = `mermaid-${Math.random().toString(36).substr(2, 9)}`;
        
        try {
          const { svg } = await mermaid.render(id, chart);
          if (containerRef.current) {
            containerRef.current.innerHTML = svg;
          }
        } catch (error) {
          console.error("Erro ao renderizar Mermaid:", error);
          containerRef.current.innerHTML = "<p class='text-red-500'>Erro na sintaxe do diagrama</p>";
        }
      }
    };

    renderChart();
  }, [chart, isMounted]);

  // Ícones SVG simples para os botões
  const IconZoomIn = () => <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6"><path strokeLinecap="round" strokeLinejoin="round" d="m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607ZM10.5 7.5v6m3-3h-6" /></svg>;
  const IconZoomOut = () => <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6"><path strokeLinecap="round" strokeLinejoin="round" d="m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607ZM13.5 10.5h-6" /></svg>;
  const IconReset = () => <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6"><path strokeLinecap="round" strokeLinejoin="round" d="M3.75 3.75v4.5m0-4.5h4.5m-4.5 0L9 9M3.75 20.25v-4.5m0 4.5h4.5m-4.5 0L9 15M20.25 3.75h-4.5m4.5 0v4.5m0-4.5L15 9m5.25 11.25h-4.5m4.5 0v-4.5m0 4.5L15 15" /></svg>;

  if (!isMounted) return <div className="text-gray-400">Carregando visualizador...</div>;

  return (
    <div className="relative w-full h-[80vh] border border-slate-700 rounded-xl overflow-hidden bg-slate-900 shadow-2xl">
      <TransformWrapper
        ref={transformComponentRef}
        initialScale={0.5} // Começa com zoom afastado para ver mais
        minScale={0.1}
        maxScale={4}
        centerOnInit={true}
        limitToBounds={false} // Permite arrastar livremente
      >
        {({ zoomIn, zoomOut, resetTransform }) => (
          <>
            {/* Barra de Ferramentas Flutuante */}
            <div className="absolute top-4 right-4 z-50 flex flex-col gap-2 bg-slate-800 p-2 rounded-lg shadow-lg border border-slate-700">
              <button onClick={() => zoomIn()} className="p-2 hover:bg-slate-700 rounded text-blue-400 transition" title="Zoom In">
                <IconZoomIn />
              </button>
              <button onClick={() => zoomOut()} className="p-2 hover:bg-slate-700 rounded text-blue-400 transition" title="Zoom Out">
                <IconZoomOut />
              </button>
              <button onClick={() => resetTransform()} className="p-2 hover:bg-slate-700 rounded text-blue-400 transition" title="Resetar Visualização">
                <IconReset />
              </button>
            </div>

            {/* Área do Diagrama */}
            <TransformComponent
              wrapperClass="w-full h-full cursor-grab active:cursor-grabbing"
              contentClass="w-full h-full"
            >
              <div 
                ref={containerRef} 
                className="mermaid-output"
                style={{ width: '100%', height: '100%' }} // Garante que o container ocupe espaço
              />
            </TransformComponent>
          </>
        )}
      </TransformWrapper>
      
      {/* Nota de rodapé (Minimap simulado textual) */}
      <div className="absolute bottom-4 left-4 z-50 bg-slate-800/80 px-3 py-1 rounded text-xs text-slate-400 pointer-events-none">
        Segure e arraste para navegar • Scroll para Zoom
      </div>
    </div>
  );
};

export default MermaidDiagram;