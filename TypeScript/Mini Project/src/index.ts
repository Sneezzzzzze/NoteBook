// Block
import crypto from 'crypto';

interface BlockShape {
    hash :string;
    prevHash :string;
    height :number;
    data :string;
}

class Block implements BlockShape {
    public hash: string;

    constructor(
        public prevHash: string,
        public height: number,
        public data: string
    ) {
        this.hash = Block.calculateHash(prevHash, height, data);
    }
    static calculateHash (
        prevHash :string,
        height :number,
        data :string
    ) {
        const toHash = `${prevHash}${height}${data}`;
        return crypto.createHash("sha256").update(toHash).digest("hex");
    }
}

class BlockChain {
    private blocks: Block[];

    constructor() {
        this.blocks = [];
    }
    private getPrevHash(): string {
        if (this.blocks.length === 0) return "";
        return this.blocks[this.blocks.length - 1].hash;
    }
    public addBlocks(data: string): void {
        const newBlock = new Block(this.getPrevHash(), this.blocks.length+1, data);
        this.blocks.push(newBlock);
    }
    public getBlocks(){
        return [...this.blocks];
    }
}

const blockChain = new BlockChain();

blockChain.addBlocks("First Block");
blockChain.addBlocks("Second Block");
blockChain.addBlocks("Third Block");

console.log(blockChain.getBlocks());