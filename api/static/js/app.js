if ('serviceWorker' in navigator) {
    window.addEventListener('load', () => {
        navigator.serviceWorker.register('/static/js/sw.js')
            .then(reg => console.log('✅ Sentinel Shield Activo'))
            .catch(err => console.error('❌ Error en el Guardián:', err));
    });
}

// Lógica para detectar si estamos offline y avisar al usuario
window.addEventListener('offline', () => {
  const status = document.getElementById('system-metrics');
  if(status) status.innerHTML = "📡 MODO OFFLINE: Usando datos locales.";
});

document.body.addEventListener('htmx:beforeRequest', function(evt) {
    console.log("🛰️ Enviando señal al Sentinel para el año: " + evt.detail.pathInfo.requestPath);
});

document.body.addEventListener('htmx:afterOnLoad', function(evt) {
    console.log("✅ Datos recibidos y renderizados en 1:1");
});

document.body.addEventListener('htmx:afterRequest', function(evt) {
    console.log("🚀 Sentinel Sync:", {
        endpoint: evt.detail.pathInfo.requestPath,
        status: evt.detail.xhr.status,
        timestamp: new Date().toLocaleTimeString()
    });
});