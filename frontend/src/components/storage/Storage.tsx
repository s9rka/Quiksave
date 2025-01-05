import { Notes } from "./MyNotes";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { ScrollArea } from "@/components/ui/scroll-area";
import watermark from '@/assets/watermark-logo.svg'

const Storage = () => {
  const userData = localStorage.getItem("userAtom");
  const username = userData ? JSON.parse(userData).username : "Guest";

  return (
    <div className="h-screen flex flex-col overflow-hidden">
      {/* Card Section */}
      <Card className="flex-grow border-none shadow-none overflow-hidden rounded-none bg-background">
        <CardHeader className="pb-3">
          <CardTitle className="text-md font-semibold flex flex-col items-start gap-2">
            {username}'s note repository
            <img className="h-8" src={watermark} />
          </CardTitle>
        </CardHeader>
        <CardContent className="h-full p-0">
          <ScrollArea className="h-full w-full">
            <div className="p-4 pb-36">
              <Notes />
            </div>
          </ScrollArea>
        </CardContent>
      </Card>

    </div>
  );
};

export default Storage;
