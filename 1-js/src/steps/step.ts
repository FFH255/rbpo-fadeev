export type Step = {
  (): Promise<void>
}

export type Steps = Step[]