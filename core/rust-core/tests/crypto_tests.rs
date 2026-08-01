use rust_core::CryptoEngine;

#[test]
fn encrypt_decrypt_roundtrip() {
    let master_key = CryptoEngine::generate_master_key();
    let engine = CryptoEngine::new(&master_key).expect("Failed to create CryptoEngine");
    let data = b"hello world";
    let (ciphertext, nonce) = engine.encrypt(data).expect("encryption failed");
    let plaintext = engine.decrypt(&ciphertext, &nonce).expect("decryption failed");
    assert_eq!(plaintext, data);
}
