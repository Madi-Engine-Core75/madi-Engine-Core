use std::io::{self, Write};
use std::net::{TcpListener, TcpStream};
use std::{thread, time};

fn handle_client(mut stream: TcpStream, addr: std::net::SocketAddr) {
    println!("🔗 [INFO] New client connected: {}", addr);
    
    let write_timeout = time::Duration::from_secs(5);
    let _ = stream.set_write_timeout(Some(write_timeout));

    loop {
        let payload = format!(
            r#"{{"status":"ONLINE","system_load":45.2,"temperature":38.4,"vortex_rpm":1200.0,"active_threads":8}}|^"#
        );

        if let Err(e) = stream.write_all(payload.as_bytes()) {
            eprintln!("❌ [ERROR] Failed to send data to client {}: {}", addr, e);
            break;
        }

        thread::sleep(time::Duration::from_secs(1));
    }
    
    println!("🔌 [INFO] Connection closed for client: {}", addr);
}

fn main() -> io::Result<()> {
    let addr = "127.0.0.1:50051";
    let listener = TcpListener::bind(addr)?;
    
    println!("🚀 Madi Engine Core (Pure Rust) is running on {}", addr);

    for stream in listener.incoming() {
        match stream {
            Ok(stream) => {
                if let Ok(client_addr) = stream.peer_addr() {
                    thread::spawn(move || {
                        handle_client(stream, client_addr);
                    });
                }
            }
            Err(e) => {
                eprintln!("❌ [ERROR] Connection failed: {}", e);
            }
        }
    }
    Ok(())
}
