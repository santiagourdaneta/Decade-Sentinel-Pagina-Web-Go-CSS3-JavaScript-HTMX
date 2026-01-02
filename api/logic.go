package main

// GetHito devuelve el hito tecnológico de un año específico.
// Esta función está aislada para poder testearla sin levantar un servidor.
func GetHito(year int) string {
	details := map[int]string{
		2016: "Auge de Docker y Microservicios. .NET Core se lanza como Open Source.",
		2017: "Publicación del paper 'Attention is All You Need'. Nace la arquitectura Transformer.",
		2018: "Lanzamiento de React Hooks. El estándar ES2018 revoluciona JavaScript.",
		2019: "HTTP/3 basado en QUIC empieza a estandarizarse. El Edge Computing despega.",
		2020: "Aceleración masiva de Cloud. Apple silicon (M1) cambia la arquitectura de CPUs portátiles.",
		2021: "Consolidación de eBPF para observabilidad. GitHub Copilot introduce la IA al desarrollo.",
		2022: "Aparición de ChatGPT y modelos estables. La ingeniería de prompts se vuelve un rol.",
		2023: "Boom de bases de datos vectoriales. HTMX revive el server-driven UI.",
		2024: "IA Agentic: Los modelos ya no solo escriben, sino que ejecutan tareas complejas.",
		2025: "Soberanía de Datos: Privacidad por diseño (GDPR 2.0) y Edge AI local.",
		2026: "Sistemas Autónomos: El software se adapta al humano, Zero-trust real.",
	}

	if val, ok := details[year]; ok {
		return val
	}
	return "Año fuera del rango de la década 2016-2026."
}