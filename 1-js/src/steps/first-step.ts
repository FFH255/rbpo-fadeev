import { getDiskInfo } from 'node-disk-info';

export async function firstStep() {
  const disks = await getDiskInfo()

  disks.forEach(disk => {
    console.log(`Диск: ${disk.filesystem}`);
    console.log(`Имя: ${disk.mounted}`);
    console.log(`Размер: ${disk.blocks / (1024 ** 3)} ГБ`);
    console.log(`Свободное место: ${disk.available / (1024 ** 3)} ГБ`);
    console.log(`Тип файловой системы: ${disk.filesystem}`);
    console.log('---');
});
}