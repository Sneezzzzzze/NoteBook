### Classes

```typescript
        // class Player {
    //     constructor(
    //         private firstName :string,
    //         private lastName :string,
    //         public nickName :string
    //     ) {}
    // }
    
    
    abstract class User {
        constructor(
            // private, protected, public
            // private : only this class can access
            // protected : only this class and its child class can access
            // public : everyone can access
            private firstName :string,
            private lastName :string,
            protected nickName :string
        )   {}
        abstract getNickName() :void
        getFullName() {
            return `${this.firstName} ${this.lastName}`;
        }
    }
    
    class Player extends User {
        getNickName() {
            console.log(this.nickName);
        }
    }
    
    const junbeom = new Player('Junbeom', 'Kim', 'Sneezzzzzze');
    // firstName, lastName is private, so you can't access them
    junbeom.nickName; // Sneezzzzzze
    junbeom.getNickName()
    
    type Words = {
        [key :string] :string;
        // key is string, value is string -> key: value
    }
    class Dict {
        private words :Words
        constructor() {
            this.words = {};
        }
    
        add(word :Word) { // Class Can be a type
            if (this.words[word.term] === undefined) {
                this.words[word.term] = word.def;
            }
        }
        update(term :string, def :string) {
            if (this.words[term] === undefined) {
                console.log('The word is not exist');
            }
            this.words[term] = def;
        }
        delete(term :string) {
            if (this.words[term] === undefined) {
                console.log('The word is not exist');
            }
            delete this.words[term];
        }
        def(term :string) {
            return this.words[term];
        }
    }
    
    class Word {
        constructor(
            public readonly term :string,
            public readonly def :string
        ) {}
    }
    
    const kimchi = new Word('kimchi', 'The best food');
    
    const dict = new Dict();
    dict.add(kimchi);
    console.log(dict.def('kimchi')); // The best food
    dict.update('kimchi', 'The best food ever');
    console.log(dict.def('kimchi')); // The best food ever  
```
