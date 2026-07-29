use rust_core::CryptoEngine;

fn main() {
    println!("Initializing Madi-Engine-Core Security Module...");

    // تهيئة محرك التشفير
    let engine = CryptoEngine::new();

    // بيانات تجريبية لتشفيرها
    let secret_data = b"Hello, this is secure vault data protected by AES-GCM-256!";
    println!("Original Data: {:?}", String::from_utf8_lossy(secret_data));

    // عملية التشفير
    match engine.encrypt(secret_data) {
        Ok((ciphertext, nonce)) => {
            println!("Encryption successful!");
            println!("Ciphertext (hex): {:x?}", ciphertext);

            // عملية فك التشفير
            match engine.decrypt(&ciphertext, &nonce) {
                Ok(decrypted_data) => {
                    println!("Decryption successful!");
                    println!("Decrypted Data: {:?}", String::from_utf8_lossy(&decrypted_data));
                }
                Err(e) => eprintln!("Decryption failed: {}", e),
            }
        }
        Err(e) => eprintln!("Encryption failed: {}", e),
    }
}
