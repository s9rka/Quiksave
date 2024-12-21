import { useParams } from "react-router-dom";
import { Notes } from "./Notes";

const Dashboard = () => {
  const { username } = useParams();

  return (
    <div>
      Welcome to the dashboard, {username}
      <h1>My notes</h1>
      <Notes />
    </div>
  );
};

export default Dashboard;
