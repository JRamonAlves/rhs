import { MoonIcon, SunIcon } from "@phosphor-icons/react"

import { Button } from "@/components/ui/button"
import { useTheme } from "./use-theme"

export function ThemeToggle() {
  const { theme, setTheme } = useTheme()
  const isDark =
    theme === "dark" ||
    (theme === "system" &&
      window.matchMedia("(prefers-color-scheme: dark)").matches)

  return (
    <Button
      variant="outline"
      size="icon"
      onClick={() => setTheme(isDark ? "light" : "dark")}
      aria-label={`Switch to ${isDark ? "light" : "dark"} theme`}
      title={`Switch to ${isDark ? "light" : "dark"} theme`}
    >
      {isDark ? <SunIcon weight="duotone" /> : <MoonIcon weight="duotone" />}
    </Button>
  )
}
