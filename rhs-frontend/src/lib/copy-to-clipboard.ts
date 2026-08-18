export async function copyToClipboard(value: string): Promise<void> {
  try {
    await navigator.clipboard.writeText(value)
    return
  } catch {
    const textarea = document.createElement("textarea")
    textarea.value = value
    textarea.readOnly = true
    textarea.style.position = "fixed"
    textarea.style.opacity = "0"
    document.body.appendChild(textarea)
    textarea.select()
    textarea.setSelectionRange(0, value.length)

    let copied = false
    try {
      copied = document.execCommand("copy")
    } finally {
      textarea.remove()
    }

    if (!copied) {
      throw new Error("Unable to copy to clipboard")
    }
  }
}
