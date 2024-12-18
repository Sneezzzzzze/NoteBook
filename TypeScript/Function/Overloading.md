### Overloading
```typescript
    // Overloading
    // when function all multiple, unique call signatures
    type Add = {
        (a: number, b: number) : number
        (a: number, b: string) : number
    }
    
    const add: Add = (a, b) => {
        a + b // error can't number + string
        
        // Not Make Sense
        if (typeof b === 'string') return a
        return a + b 
    }
    type Config = {
        path: string,
        state: object
    }
    type Push = {
        (path:string): void
        (config: Config): void
    }
    const push:Push = (config) => {
        if (typeof config === 'string') console.log(config)
        else {
            console.log(config.path, config.state)
        }
    }
    // if diffenent amount
    type Add = {
        (a:number, b:number) :number
        (a:number, b:number, c:number) :number
    }
    const add:Add = (a, b, c?:number) => { // need to defined it to number
        if(c) return a + b + c
        return a + b
    }
    add(1, 2) // 3
    add(1, 2, 3) // 6
    
```