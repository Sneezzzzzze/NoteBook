### Polymorphism

```typescript
    // Poly means many, much ,multi
    // morphos means form, structure
    // <> -> Generic ( use this to let TypeScript find type with itselfs )
    type SuperPrint = {
        (arr :number[]) :void
        (arr :boolean[]) :void
        (arr :string[]) :void
        // change to 
        <T>(arr :T[]) :T
    }
    
    const superPrint: SuperPrint = (arr) => arr[0]
    superPrint([1, true, 'hee']) // it will work
```