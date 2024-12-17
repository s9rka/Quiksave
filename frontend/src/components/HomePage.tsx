import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";

const HomePage = () => {
  return (
    <div>
      <Card>
        <CardHeader>
          <CardTitle>nota-bene</CardTitle>
          <CardDescription>Simple powerful notepad</CardDescription>
        </CardHeader>
        <CardContent>
          <p>Love 📜</p>
        </CardContent>
        <CardFooter>
          <p>Sorender Solutions Unlimited</p>
        </CardFooter>
      </Card>
    </div>
  );
};

export default HomePage;
