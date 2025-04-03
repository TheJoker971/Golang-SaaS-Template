import { Button } from "../components/ui/button"
import { Container } from "../components/ui/container"
import { Logo } from "../components/ui/logo"
import { Link } from "react-router-dom"

export function SiteHeader() {
  return (
    <header className="border-b py-4">
      <Container>
        <div className="flex items-center justify-between">
          <Logo />

          <div className="space-x-4">
            <Link to="/login">
              <Button variant="outline">Connexion</Button>
            </Link>
            <Link to="/register">
              <Button>Inscription</Button>
            </Link>
          </div>
        </div>
      </Container>
    </header>
  )
}
