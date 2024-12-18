### Call Signatures

```typescript
    // Call Signatures
    type Add = (a:number, b:number) => number;

    const add:Add = (a, b) => a + b;
    const add:Add = (a, b) => {a + b}; // If add { } It Gona me void
```
