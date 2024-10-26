import * as fs from 'fs/promises';
import { prompt } from "../io";

export async function secondStep() {
  const fileName = 'userInput.txt';

  // 1. Запросить строку у пользователя
  const userInput = await prompt("Введите строку для записи в файл: ");

  // 2. Создать файл и записать строку
  await fs.writeFile(fileName, userInput, 'utf-8');
  console.log(`Строка записана в файл ${fileName}`);

  // 3. Прочитать содержимое файла
  const fileContent = await fs.readFile(fileName, 'utf-8');
  console.log(`Содержимое файла ${fileName}:`);
  console.log(fileContent);

  // 4. Удалить файл
  await fs.unlink(fileName);
  console.log(`Файл ${fileName} удалён.`);
}