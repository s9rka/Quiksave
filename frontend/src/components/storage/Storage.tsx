import { useParams } from "react-router-dom";
import { Notes } from "./Notes";

const Dashboard = () => {
  const { username } = useParams();

  return (
    <div>
      <h1>{username}'s storage</h1>
      <Notes />
    </div>
  );
};

export default Dashboard;
