import React from "react";
import { toObjectArray } from "../utils/utils";
import { IData } from "../types/data";
import AWS from "aws-sdk";

interface FileImporterProps {
  onImportFile: (data: IData) => void;
}

export const FileImporter = ({ onImportFile }: FileImporterProps) => {
  const saveTransactions = async (jsonData: string) => {
    try {
      const result = await fetch(
        `${process.env.REACT_APP_API_GATEWAY_ENDPOINT}transactions`,
        {
          method: "POST",
          body: jsonData,
        }
      );

      console.log("Result", await result.json());
    } catch (e: any) {
      console.log("Error: ==> ", e.message);
    }
  };

  const uploadFileToS3 = async (file: File): Promise<string | null> => {
    const fileName = `summary-${Date.now()}-${file?.name}`;

    AWS.config.update({
      accessKeyId: process.env.REACT_APP_AWS_S3_ACCESS_KEY_ID,
      secretAccessKey: process.env.REACT_APP_AWS_S3_SECRET_ACCESS_KEY,
    });

    const s3 = new AWS.S3({
      params: {
        Bucket: process.env.REACT_APP_AWS_S3_BUCKET,
        region: process.env.REACT_APP_AWS_REGION,
      },
    });

    try {
      const params = {
        Bucket: process.env.REACT_APP_AWS_S3_BUCKET,
        Key: fileName,
        Body: file,
      };

      await s3
        // @ts-ignore
        .putObject(params)
        .on("httpUploadProgress", (evt) => {
          console.log(
            // @ts-ignore
            "Uploading " + parseInt((evt.loaded * 100) / evt.total) + "%"
          );
        })
        .promise();

      return fileName;
    } catch (error) {
      console.error("Error uploading file:", error);
      return null;
      // Handle errors gracefully, informing the user if needed
    }
  };

  const handleUpload = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files ? e.target.files[0] : null;

    let jsonData = "";

    // Read CSV file
    const reader = new FileReader();
    reader.onload = async () => {
      // @ts-ignore
      const rows = reader.result.split("\n");
      const parsedData = rows.map((row: any) => row.split(","));

      // setFile(file); // Set file info

      // Set headers and data to show table in screen
      onImportFile({
        headers: parsedData[0],
        data: parsedData.slice(1),
      });

      // Convert JSON to object array
      const objectArray = toObjectArray(parsedData.slice(1));
      jsonData = JSON.stringify(objectArray);

      saveTransactions(jsonData);
    };

    if (file) {
      uploadFileToS3(file);
      reader.readAsText(file);
    }
  };

  return (
    <div className="text-center">
      <p className="text-base font-semibold text-indigo-600">
        Transactions File Manager
      </p>
      <h1 className="mt-4 text-3xl font-bold tracking-tight text-gray-900 sm:text-5xl">
        Get Your Transaction Summary in a Snap!
      </h1>
      <p className="mt-6 text-base leading-7 text-gray-600">
        Just provide your email and a valid CSV file, and we'll deliver a clear
        overview directly to your inbox.
      </p>

      <div className="w-full md:w-2/5 mx-auto mt-10">
        <label
          htmlFor="email-input"
          className="block mb-2 text-md font-medium text-gray-900"
        >
          Your email
        </label>
        <input
          type="email"
          required
          id="email-input"
          aria-describedby="helper-text-explanation"
          className="border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 font-bold"
          placeholder="john@doe.com"
        />
        <p id="helper-text-explanation" className="mt-2 text-sm text-gray-500">
          We’ll never share your details. Read our{" "}
          <a href="#" className="font-medium text-blue-600 hover:underline">
            Privacy Policy
          </a>
          .
        </p>
      </div>

      <div className="mt-10 flex items-center justify-center gap-x-6">
        <label htmlFor="dropzone-file">
          <div className="cursor-pointer py-2 px-4 flex justify-center items-center bg-red-600 hover:bg-red-700 focus:ring-red-500 focus:ring-offset-red-200 text-white w-full transition ease-in duration-200 text-center text-base font-semibold shadow-md focus:outline-none focus:ring-2 focus:ring-offset-2 rounded-lg">
            <svg
              width="20"
              height="20"
              fill="currentColor"
              className="mr-2"
              viewBox="0 0 1792 1792"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path d="M1344 1472q0-26-19-45t-45-19-45 19-19 45 19 45 45 19 45-19 19-45zm256 0q0-26-19-45t-45-19-45 19-19 45 19 45 45 19 45-19 19-45zm128-224v320q0 40-28 68t-68 28h-1472q-40 0-68-28t-28-68v-320q0-40 28-68t68-28h427q21 56 70.5 92t110.5 36h256q61 0 110.5-36t70.5-92h427q40 0 68 28t28 68zm-325-648q-17 40-59 40h-256v448q0 26-19 45t-45 19h-256q-26 0-45-19t-19-45v-448h-256q-42 0-59-40-17-39 14-69l448-448q18-19 45-19t45 19l448 448q31 30 14 69z"></path>
            </svg>
            Upload your .csv file
          </div>
        </label>

        <input
          id="dropzone-file"
          accept="text/csv"
          type="file"
          className="hidden"
          onChange={handleUpload}
        />
      </div>
    </div>
  );
};
