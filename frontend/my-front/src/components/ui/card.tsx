
import { ReactNode } from "react"
import { cn } from "../../lib/utils"

export function Card({ children, className }: { children: ReactNode; className?: string }) {
  return <div className={cn("bg-white shadow rounded-lg", className)}>{children}</div>
}

export function CardHeader({ children }: { children: ReactNode }) {
  return <div className="border-b px-4 py-2">{children}</div>
}

export function CardTitle({ children }: { children: ReactNode }) {
  return <h2 className="text-lg font-semibold">{children}</h2>
}

export function CardDescription({ children }: { children: ReactNode }) {
  return <p className="text-sm text-gray-500">{children}</p>
}

export function CardContent({ children }: { children: ReactNode }) {
  return <div className="p-4">{children}</div>
}

export function CardFooter({ children }: { children: ReactNode }) {
  return <div className="border-t px-4 py-2">{children}</div>
}
