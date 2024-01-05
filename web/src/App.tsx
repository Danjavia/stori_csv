import React, { useEffect, useState } from "react";
import "./App.css";
import { TransactionList } from "./components/transactionList";
import { FileImporter } from "./components/fileImporter";
import { IData } from "./types/data";
import Logo from "./assets/images/logo.svg";
import { useKindeAuth } from "@kinde-oss/kinde-auth-react";

const App = () => {
  const { login, logout, isAuthenticated, user, getUser } = useKindeAuth();
  const [data, setData] = useState<string[][] | null>(null);
  const [headers, setHeaders] = useState<string[] | null>(null);

  const onImportFile = (data: IData) => {
    setHeaders(data.headers);
    setData(data.data);
  };

  const signin = () => {
    login();
  };

  useEffect(() => {
    console.log(user);
    const u = getUser;
    console.log(u);
  }, []);

  return (
    <div>
      <header className="app-header flex justify-between items-center p-4">
        <img src={Logo} alt="Logo" />
        <div>
          {isAuthenticated ? (
            <button
              onClick={logout}
              className="w-auto cursor-pointer inline-block py-2 px-4 justify-center items-center transition ease-in duration-200 text-center text-base font-semibold focus:outline-none focus:ring-2 focus:ring-offset-2 border-b-2 hover:border-b-lime-900 text-lime-900"
            >
              Login
            </button>
          ) : (
            <button
              onClick={signin}
              className="w-auto cursor-pointer inline-block py-2 px-4 justify-center items-center transition ease-in duration-200 text-center text-base font-semibold focus:outline-none focus:ring-2 focus:ring-offset-2 border-b-2 hover:border-b-lime-900 text-lime-900"
            >
              Login
            </button>
          )}
        </div>
      </header>
      <main className="grid min-h-screen place-items-center bg-white px-6 py-12 sm:py-12 lg:px-8">
        {!headers && <FileImporter onImportFile={onImportFile} />}

        {headers && data && <TransactionList headers={headers} data={data} />}

        <div className="text-center">
          <h2 className="font-bold text-lg">Powered By</h2>
          <img src={Logo} alt="Stori" className="mb-4" />
        </div>
      </main>

      {/*{file && <pre>{jsonData}</pre>}*/}
    </div>
  );
};

export default App;
