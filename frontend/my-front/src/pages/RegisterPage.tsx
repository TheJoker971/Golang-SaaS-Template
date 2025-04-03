
import { useState } from 'react'

export default function RegisterPage() {
  const [form, setForm] = useState({ pseudo: '', email: '', password: '', admin: false })

  const handleSubmit = async (e) => {
    e.preventDefault()
    const res = await fetch('http://localhost:8080/register', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(form)
    })

    if (res.status === 201) {
      alert('Inscription réussie')
    } else {
      const err = await res.text()
      alert('Erreur : ' + err)
    }
  }

  return (
    <form onSubmit={handleSubmit}>
      <h2>Inscription</h2>
      <input type="text" placeholder="Pseudo" value={form.pseudo} onChange={(e) => setForm({ ...form, pseudo: e.target.value })} />
      <input type="email" placeholder="Email" value={form.email} onChange={(e) => setForm({ ...form, email: e.target.value })} />
      <input type="password" placeholder="Mot de passe" value={form.password} onChange={(e) => setForm({ ...form, password: e.target.value })} />
      <button type="submit">S'inscrire</button>
    </form>
  )
}
