
import { InputHTMLAttributes } from "react"

export function Checkbox(props: InputHTMLAttributes<HTMLInputElement>) {
  return <input type="checkbox" className="h-4 w-4 text-blue-600 rounded" {...props} />
}
