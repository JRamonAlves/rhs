import { StrictMode } from "react"
import { createRoot } from "react-dom/client"

import { ThemeProvider } from "@/components/theme"
import App from "./App"
import "./index.css"
import { QueryClient, QueryClientProvider } from "@tanstack/react-query"

const rootElement = document.getElementById("root")

if (!rootElement) {
  throw new Error("Root element not found")
}

const queryClient = new QueryClient()

createRoot(rootElement).render(
  <StrictMode>
    <ThemeProvider>
      <QueryClientProvider client={queryClient}>
        <App />
      </QueryClientProvider>
    </ThemeProvider>
  </StrictMode>
)
