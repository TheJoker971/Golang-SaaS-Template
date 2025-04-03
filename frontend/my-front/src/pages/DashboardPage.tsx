import { Button } from "../components/ui/button"
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "../components/ui/card"
import { SiteHeader } from "../components/site-header"
import {
  BarChart3,
  Bell,
  Calendar,
  CheckCircle,
  FileText,
  MessageSquare,
  Plus,
  Settings,
  User,
  Users,
} from "lucide-react"
import { Link } from "react-router-dom"
import "./dashboard.css"

export default function DashboardPage() {
  const username = "Utilisateur"
  const stats = [
    { title: "Projets", value: "12", icon: <FileText /> },
    { title: "Tâches", value: "42", icon: <CheckCircle /> },
    { title: "Messages", value: "8", icon: <MessageSquare /> },
    { title: "Réunions", value: "3", icon: <Calendar /> },
  ]

  const recentActivities = [
    { title: "Projet mis à jour", time: "Il y a 2 heures", description: "Vous avez modifié le projet 'Marketing Q2'" },
    {
      title: "Nouveau message",
      time: "Il y a 5 heures",
      description: "Marie vous a envoyé un message concernant la réunion",
    },
    { title: "Tâche terminée", time: "Hier", description: "Vous avez terminé la tâche 'Préparer la présentation'" },
  ]

  return (
    <div className="dashboard">
      <SiteHeader />

      <main className="main">
        <div className="container">
          {/* Welcome banner */}
          <div className="welcome-banner card">
            <div className="banner-content">
              <div className="text">
                <h1>Bienvenue, {username}!</h1>
                <p>Vous vous êtes connecté avec succès. Voici un aperçu de votre tableau de bord.</p>
                <div className="actions">
                  <Button>
                    <Plus /> Nouveau projet
                  </Button>
                  <Button className="outline">
                    <BarChart3 /> Voir les statistiques
                  </Button>
                </div>
              </div>
              <div className="image">
                <img src="https://placehold.co/200x200" alt="Dashboard illustration" />
              </div>
            </div>
          </div>

          {/* Stat cards */}
          <div className="stats grid">
            {stats.map((stat, index) => (
              <Card key={index} className="card">
                <CardHeader className="card-header">
                  <CardTitle>{stat.title}</CardTitle>
                  {stat.icon}
                </CardHeader>
                <CardContent>
                  <div className="stat-value">{stat.value}</div>
                </CardContent>
              </Card>
            ))}
          </div>

          <div className="dashboard-grid">
            {/* Recent activities */}
            <Card className="card">
              <CardHeader>
                <CardTitle>Activité récente</CardTitle>
                <CardDescription>Voici vos dernières activités sur la plateforme</CardDescription>
              </CardHeader>
              <CardContent>
                {recentActivities.map((activity, index) => (
                  <div key={index} className="activity">
                    <div className="icon-box">
                      <Bell />
                    </div>
                    <div>
                      <h4>{activity.title}</h4>
                      <p>{activity.description}</p>
                      <p className="time">{activity.time}</p>
                    </div>
                  </div>
                ))}
              </CardContent>
              <CardFooter>
                <Button className="outline full">Voir toutes les activités</Button>
              </CardFooter>
            </Card>

            {/* Quick links */}
            <Card className="card">
              <CardHeader>
                <CardTitle>Accès rapides</CardTitle>
                <CardDescription>Liens utiles pour naviguer dans l'application</CardDescription>
              </CardHeader>
              <CardContent>
                <Button className="ghost full"><User /> Mon profil</Button>
                <Button className="ghost full"><Users /> Mon équipe</Button>
                <Button className="ghost full"><FileText /> Mes projets</Button>
                <Button className="ghost full"><Calendar /> Calendrier</Button>
                <Button className="ghost full"><Settings /> Paramètres</Button>
              </CardContent>
              <CardFooter>
                <Button className="outline full">Centre d'aide</Button>
              </CardFooter>
            </Card>
          </div>
        </div>
      </main>

      <footer className="footer">
        <div className="container footer-content">
          <p>© {new Date().getFullYear()} VotreEntreprise. Tous droits réservés.</p>
          <div className="footer-links">
            <Link to="#">Confidentialité</Link>
            <Link to="#">Conditions</Link>
            <Link to="#">Contact</Link>
          </div>
        </div>
      </footer>
    </div>
  )
}
