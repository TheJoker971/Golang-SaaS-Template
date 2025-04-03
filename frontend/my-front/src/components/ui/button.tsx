
import { cn } from "../../lib/utils"
import { ButtonHTMLAttributes, forwardRef } from "react"

export interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: "primary" | "secondary" | "outline" | "ghost"
}

export const Button = forwardRef<HTMLButtonElement, ButtonProps>(
  ({ className, variant = "primary", ...props }, ref) => {
    const base = "px-4 py-2 rounded font-medium"

    const variants = {
      primary: "bg-blue-600 text-white hover:bg-blue-700",
      secondary: "bg-gray-100 text-black hover:bg-gray-200",
      outline: "border border-gray-300 text-black bg-white hover:bg-gray-50",
      ghost: "bg-transparent text-black hover:bg-gray-100",
    }

    return (
      <button
        ref={ref}
        className={cn(base, variants[variant], className)}
        {...props}
      />
    )
  }
)

Button.displayName = "Button"
