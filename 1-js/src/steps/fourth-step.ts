import * as fs from 'fs/promises';
import { prompt } from "../io";

export async function fourthStep() {
  const fileName = 'userInput.xml';

  const userName = await prompt("Введите имя: ");
  const userAge = await prompt("Введите возраст: ");

  const user = `
    <?xml version="1.0" encoding="UTF-8"?>
    <user>
      <name>${userName}</name>
      <age>${userAge}</age>
    </user>
  `;

  await fs.writeFile(fileName, user, 'utf-8');
  console.log(`XML записан в файл ${fileName}`);

  const fileContent = await fs.readFile(fileName, 'utf-8');
  console.log(`Содержимое файла ${fileName}:`);
  console.log(fileContent);

  await fs.unlink(fileName);
  console.log(`Файл ${fileName} удалён.`);
}