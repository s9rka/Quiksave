import { Notes } from "./MyNotes";
import { Card, CardContent, CardTitle } from "@/components/ui/card";
import watermark from "@/assets/watermark-logo.svg";

const Storage = () => {
  const userData = localStorage.getItem("userAtom");
  const username = userData ? JSON.parse(userData).username : "Guest";

  return (
    <div className="py-10 flex flex-col overflow-hidden">
      <Card className="px-4 flex-grow border-none shadow-none overflow-hidden rounded-none bg-background">
        <img className="h-6 self-start mb-2" src={watermark} />
        <CardTitle className="font-normal text-xl">
          {username}'s note repository
        </CardTitle>
        <CardContent className="h-full p-0">
          <Notes />
        </CardContent>
      </Card>
    </div>
  );
};

export default Storage;
