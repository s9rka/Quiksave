import { Notes } from "./MyNotes";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import watermark from "@/assets/watermark-logo.svg";
import { useVault as useVaultQuery } from "@/services/queries";
import { useParams } from "react-router-dom";


const Storage = () => {
  const userData = localStorage.getItem("userAtom");
  const username = userData ? JSON.parse(userData).username : "Guest";
  const params = useParams();

  const { data: vaultData } = useVaultQuery(Number(params.id));

  

  return (
    <div className="py-10 flex flex-col overflow-hidden">
      <img className="h-6 mx-auto" src={watermark} />
      <Card className="px-4 flex-grow border-none shadow-none overflow-hidden rounded-none bg-background">
        <CardHeader className="flex flex-row justify-between ">
          <div>
            <CardTitle className="font-normal text-xl">
              {vaultData?.name}
            </CardTitle>
          </div>
          <p>{username}</p>
        </CardHeader>
        <CardContent className="h-full">
          <Notes />
        </CardContent>
      </Card>
    </div>
  );
};

export default Storage;
