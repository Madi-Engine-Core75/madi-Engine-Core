mod crypto;
use crypto::CryptoEngine;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    // توليد مفتاح رئيسي جديد للاختبار
    let master_key = CryptoEngine::generate_master_key();

    // تهيئة محرك التشفير بالمفتاح المولّد
    let crypto_engine = CryptoEngine::new(&master_key)
        .map_err(|e| format!("Failed to initialize CryptoEngine: {}", e))?;

    let plaintext = b"MadiEngineCore Secure Payload Test";
    
    // تجربة التشفير
    let (ciphertext, nonce) = crypto_engine.encrypt(plaintext)
        .map_err(|e| format!("Encryption error: {}", e))?;

    // تجربة فك التشفير
    let decrypted = crypto_engine.decrypt(&ciphertext, &nonce)
        .map_err(|e| format!("Decryption error: {}", e))?;

    assert_eq!(&plaintext[..], &decrypted[..]);
    println!("Crypto engine verification passed successfully!");

    Ok(())
}
