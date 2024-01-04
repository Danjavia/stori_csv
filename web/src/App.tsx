import React, { useState } from "react";
import "./App.css";
import { TransactionList } from "./components/transactionList";
import { FileImporter } from "./components/fileImporter";
import { IData } from "./types/data";
import Logo from "./assets/images/logo.svg";

const App = () => {
  const [data, setData] = useState<string[][] | null>(null);
  const [headers, setHeaders] = useState<string[] | null>(null);

  const onImportFile = (data: IData) => {
    setHeaders(data.headers);
    setData(data.data);
  };

  return (
    <div>
      <main className="grid min-h-screen place-items-center bg-white px-6 py-12 sm:py-12 lg:px-8">
        {!headers && <FileImporter onImportFile={onImportFile} />}

        {headers && data && <TransactionList headers={headers} data={data} />}

        <div className="text-center">
          <h2 className="font-bold text-xl">Powered By</h2>
          <img src={Logo} alt="Stori" className="mb-4" />
        </div>
      </main>

      {/*{file && <pre>{jsonData}</pre>}*/}
    </div>
  );
};

export default App;
