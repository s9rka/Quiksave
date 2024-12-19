import { useParams } from "react-router-dom"

const Dashboard = () => {
    const {username} = useParams()

  return (
    <div>
      Welcome to the dashboard, {username}
    </div>
  )
}

export default Dashboard
