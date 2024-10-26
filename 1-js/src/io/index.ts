import * as readline from 'readline';

export function prompt(query: string): Promise<string> {
    const rl = readline.createInterface({
        input: process.stdin,
        output: process.stdout,
    });

    return new Promise(resolve => rl.question(query, (answer) => {
        rl.close();
        resolve(answer);
    }));
}

export function pressEnterToContinue(): Promise<void> {
    const rl = readline.createInterface({
        input: process.stdin,
        output: process.stdout,
    });

    return new Promise((resolve) => {
        rl.question("Нажмите Enter, чтобы продолжить...", () => {
            rl.close();
            resolve();
        });
    });
}