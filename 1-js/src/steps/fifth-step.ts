import archiver from 'archiver';
import * as fs from 'fs';
import * as fsPromises from 'fs/promises';
import * as path from 'path';
import * as unzipper from 'unzipper';
import { prompt } from '../io';

const zipFileName = 'archive.zip';

export async function fifthStep() {
  // 1. Создать ZIP-архив
  const output = fs.createWriteStream(zipFileName);
  const archive = archiver('zip', { zlib: { level: 9 } });

  // Обработать события завершения
  output.on('close', () => {
      console.log(`ZIP архив ${zipFileName} создан (${archive.pointer()} байт).`);
  });

  archive.pipe(output);
  debugger;

  // 2. Запросить файл у пользователя и добавить его в архив
  const filePath = await prompt("Введите путь к файлу для добавления в архив: ");
  const fileName = path.basename(filePath);
  archive.file(filePath, { name: fileName });
  await archive.finalize();

  console.log(`Файл ${fileName} добавлен в архив ${zipFileName}.`);

  // 3. Разархивировать файл и вывести данные о нем
  const extractedDir = 'extracted';
  await fsPromises.mkdir(extractedDir, { recursive: true });

  await new Promise(resolve => {
    fs.createReadStream(zipFileName)
    .pipe(unzipper.Extract({ path: extractedDir }))
    .on('close', async () => {
        console.log(`Файл из архива разархивирован в папку ${extractedDir}.`);
        
        // Вывод информации о разархивированном файле
        const extractedFilePath = path.join(extractedDir, fileName);
        const stats = await fsPromises.stat(extractedFilePath);
        console.log(`Имя файла: ${fileName}`);
        console.log(`Размер файла: ${stats.size} байт`);
        console.log(`Дата создания: ${stats.birthtime}`);
        console.log(`Дата последнего изменения: ${stats.mtime}`);
        resolve(true)
    });
  })

  // 4. Удалить файл и архив
  await fsPromises.unlink(zipFileName);
  await fsPromises.rm(extractedDir, { recursive: true, force: true });
  console.log(`Файл ${zipFileName} и папка ${extractedDir} удалены.`);
  
}