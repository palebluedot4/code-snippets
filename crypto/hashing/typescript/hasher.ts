export interface Hasher {
  hash(data: Uint8Array): Promise<Uint8Array>;
  verify(hash: Uint8Array, data: Uint8Array): Promise<boolean>;
}
