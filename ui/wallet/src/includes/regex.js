
export const domain = new RegExp(/(?=^[^.]{1,15}\..*$)(?=^.{1,254}$)(^(?:(?!\d+\.)[a-zA-Z0-9_-]{1,63}\.?)+(?:[a-zA-Z]{2,})$)/)

export const mnemonic = new RegExp(/^\w+(?:\s\w+){11,}?/)

export const password = new RegExp(/^[a-zA-Z0-9-._!@#$%^&*()]{6,20}$/)

export const secret = new RegExp(/^[0-9]{4,6}$/)

export const username = new RegExp(/[a-zA-Z0-9-._]{3,20}/)
