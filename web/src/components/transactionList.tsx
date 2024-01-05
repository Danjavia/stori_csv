import React from "react";
import { useKindeAuth } from "@kinde-oss/kinde-auth-react";
import { clsx } from "clsx";

interface TransactionListProps {
  headers: string[];
  data: string[][];
}

export const TransactionList = ({ headers, data }: TransactionListProps) => {
  const { register } = useKindeAuth();

  const authorize = () => {
    register({
      authUrlParams: {
        email: "danjavia@gmail.com",
      },
    });
  };

  return (
    <>
      <div className="text-center mb-4">
        <h1 className="mt-4 text-3xl font-bold tracking-tight text-gray-900 sm:text-5xl">
          Got it! Your file content is ready to explore.
        </h1>
        <p className="mt-6 text-base leading-7 text-gray-600">
          Just sit back and relax! We'll quickly generate a helpful summary of
          your transactions and send it right to the email you provided.
        </p>
      </div>

      <div className="overflow-x-auto bg-white p-4 w-full md:w-1/2 rounded shadow-xl mx-auto mt-4">
        <table className="min-w-full text-left text-sm whitespace-nowrap">
          <thead className="uppercase tracking-wider border-b-2 bg-neutral-50">
            <tr>
              {headers.map((header) => (
                <th scope="col" className="px-6 py-4" key={header}>
                  {header}
                </th>
              ))}
            </tr>
          </thead>
          <tbody>
            {data.map((row, index) => (
              <tr key={index} className="border-b hover:bg-neutral-100">
                {row.map((cell, i) => {
                  const className = clsx(
                    "px-6 py-4 whitespace-nowrap",
                    cell.startsWith("-") && "font-bold text-red-600",
                    cell.startsWith("+") && "font-bold text-green-600"
                  );

                  return (
                    <td key={i} className={className}>
                      {cell}
                    </td>
                  );
                })}
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      <div className="mx-auto mt-10 text-center mb-10">
        <h4 className="font-bold text-xl text-center mb-4">
          Unlock full control! Save your data and follow your summary online for
          deeper insights.
        </h4>

        <button
          onClick={authorize}
          className="w-auto cursor-pointer inline-block py-2 px-4 justify-center items-center bg-green-600 hover:bg-green-700 focus:ring-red-500 focus:ring-offset-red-200 text-white transition ease-in duration-200 text-center text-base font-semibold shadow-md focus:outline-none focus:ring-2 focus:ring-offset-2 rounded-lg"
        >
          Create my account and track
        </button>
      </div>
    </>
  );
};
