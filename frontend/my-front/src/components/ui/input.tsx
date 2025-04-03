
import { InputHTMLAttributes, forwardRef } from "react"
import { cn } from "../../lib/utils"

export interface InputProps extends InputHTMLAttributes<HTMLInputElement> {}

export const Input = forwardRef<HTMLInputElement, InputProps>(({ className, ...props }, ref) => {
  return <input ref={ref} className={cn("border border-gray-300 rounded px-3 py-2", className)} {...props} />
})

Input.displayName = "Input"
