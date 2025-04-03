import { useState } from 'react'
import { useNavigate } from 'react-router-dom'

export default function LoginPage() {
  const [login, setLogin] = useState('')
  const [password, setPassword] = useState('')
  const navigate = useNavigate() // ← Hook de navigation

  const handleLogin = async (e) => {
    e.preventDefault()
    const res = await fetch('http://localhost:8080/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ login, password }),
    })

    if (!res.ok) {
      alert('Erreur de connexion')
      return
    }

    const user = await res.json()
    console.log(user.Pseudo)
    alert(`Connecté : ${user}`)

    // Redirection vers /dashboard
    navigate('/dashboard')
  }

  return (
    <form onSubmit={handleLogin}>
      <h2>Connexion</h2>
      <input type="text" placeholder="Email ou pseudo" value={login} onChange={(e) => setLogin(e.target.value)} />
      <input type="password" placeholder="Mot de passe" value={password} onChange={(e) => setPassword(e.target.value)} />
      <button type="submit">Se connecter</button>
    </form>
  )
}
