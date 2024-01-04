import { Transaction } from "../types/transaction";

export const toObjectArray = (data: [][]): Transaction[] => {
  return (
    data.map((row: any) => ({
      id: parseInt(row[0], 10),
      date: row[1],
      transaction: Number(row[2]),
    })) || []
  );
};
