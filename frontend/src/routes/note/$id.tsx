import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/note/$id')({
  component: RouteComponent,
})

function RouteComponent() {
    const { id } = Route.useParams()
    
  return <div>Hello /note {id}!</div>
}
