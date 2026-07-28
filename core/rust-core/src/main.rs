mod crypto;
mod auth;

use crypto::CryptoEngine;
use auth::AuthValidator;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let master_key = CryptoEngine::generate_master_key();
    let crypto_engine = CryptoEngine::new(&master_key)
        .map_err(|e| format!("Failed to initialize CryptoEngine: {}", e))?;

    // محاكاة طلب وارد من البوابة مع التوكن والبيانات المشفرة
    let incoming_token = "MADI_SECURE_TOKEN_123";
    let plaintext = b"Microservice Payload Transaction";

    // الخطوة 1: التحقق من المصادقة أولاً
    let is_authorized = AuthValidator::verify_token(incoming_token)
        .map_err(|e| format!("Auth error: {}", e))?;

    if is_authorized {
        println!("Authentication passed successfully.");

        // الخطوة 2: التشفير والمعالجة في النواة
        let (ciphertext, nonce) = crypto_engine.encrypt(plaintext)?;
        let decrypted = crypto_engine.decrypt(&ciphertext, &nonce)?;

        assert_eq!(&plaintext[..], &decrypted[..]);
        println!("Microservice pipeline verified successfully!");
    }

    Ok(())
}
