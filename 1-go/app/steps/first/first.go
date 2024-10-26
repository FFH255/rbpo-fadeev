package first

import (
	"bufio"
	"fmt"
	"github.com/rodaine/table"
	"github.com/shirou/gopsutil/disk"
)

func Step(_ *bufio.Reader) {
	tbl := table.New("Диск", "Точка монтирования", "Файловая система", "Общий размер", "Свободное место", "Использованное место")

	// Получаем информацию о всех разделах диска
	partitions, err := disk.Partitions(true)
	if err != nil {
		fmt.Println("Ошибка получения разделов диска:", err)
		return
	}

	// Проходим по каждому разделу и выводим информацию
	for _, partition := range partitions {

		// Получаем информацию о свободном и общем месте на диске
		usage, err := disk.Usage(partition.Mountpoint)
		if err == nil {
			tbl.AddRow(
				partition.Device,
				partition.Mountpoint,
				partition.Fstype,
				fmt.Sprintf("%d Гб", usage.Total/1024/1024/1024),
				fmt.Sprintf("%d Гб", usage.Free/1024/1024/1024),
				fmt.Sprintf("%f%%", usage.UsedPercent),
			)
		}
	}
	tbl.Print()
}
