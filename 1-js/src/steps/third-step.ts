import * as fs from 'fs/promises';
import { prompt } from "../io";

export async function thirdStep() {
  const fileName = 'userInput.json';

  const userName = await prompt("Введите имя: ");
  const userAge = await prompt("Введите возраст: ");

  const user = { name: userName, age: userAge };

  const data = JSON.stringify(user, null, 2);

  await fs.writeFile(fileName, data, 'utf-8');
  console.log(`JSON записан в файл ${fileName}`);

  const fileContent = await fs.readFile(fileName, 'utf-8');
  console.log(`Содержимое файла ${fileName}:`);
  console.log(fileContent);

  await fs.unlink(fileName);
  console.log(`Файл ${fileName} удалён.`);
}