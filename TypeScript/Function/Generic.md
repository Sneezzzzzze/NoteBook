### Generic
```typescript
    type SuperPrint = <T, M>(arr :T[], b :M) => T
    superPrint([1, true, "x"], "hee")

    function superPrint<V>(a :V[]) {
        return a[0]
    }
    
    type Player<E> = {
        name :string
        extraInfo :E
    }
    type junExtra = {
        favFood :string
    }
    type junPlayer = Player<junExtra>
    
    const jun :junPlayer<{favFood :string}> = {
        name :"junbeom",
        extraInfo :{
            favFood :"kimchi"
        }
    }
    const lynn :Player<null> = {
        name :"lynn",
        extraInfo : {}
            favFood :"kimchi"
    }
    
    type Arr  = Array<number>
    function printAllNumbers(arr: Array<number>) {
        ...
    }
```