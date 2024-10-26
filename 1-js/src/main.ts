import { pressEnterToContinue } from './io';
import { Steps, fifthStep, firstStep, secondStep, thirdStep } from './steps';
import { fourthStep } from './steps/fourth-step';

async function main() {
    const steps: Steps = [
        firstStep,
        secondStep,
        thirdStep,
        fourthStep,
        fifthStep,
    ]

    console.log("ПРОГРАММА ЗАПУЩЕНА\n")

    for (let i = 0; i < steps.length; i++) {
        const step = steps[i]
        console.log(`Шаг: ${i+1}\n`) 
        await step()
        console.log()
        await pressEnterToContinue()
    }

    console.log("\nПРОГРАММА ЗАВЕРШЕНА")
}

main().catch(error => console.error(error));
