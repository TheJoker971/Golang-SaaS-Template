import { Outlet } from 'react-router-dom'
import { SiteHeader } from './components/site-header'

export default function App() {
  return (
    <>
      <SiteHeader />
      <main className="p-4">
        <Outlet />
      </main>
    </>
  )
}
