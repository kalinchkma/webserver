use actix_web::{get, guard, web, App, HttpResponse, HttpServer, Responder};
use std::sync::Mutex;

struct AppState {
    app_name: String,
    counter: Mutex<i32>
}

#[get("/hello-world")]
async fn hello() -> impl Responder {
    HttpResponse::Ok().body("Hello world")
}

#[get("/")]
async fn index(data: web::Data<AppState>) -> String {
    let app_name = &data.app_name;
    let mut counter = data.counter.lock().unwrap();
    *counter += 1;
    format!("This is the cool name of all time {app_name} : Counter - {counter}")
}

async fn manual_hello() -> impl Responder {
    HttpResponse::Ok().body("Manual hello world")
}



#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
         .app_data(web::Data::new(AppState {
            app_name: String::from("Popcorn"),
            counter: Mutex::new(0)
         })).service(index)
         .service(hello)
         .service(
            web::scope("/api")
            .route("/manual", web::get().to(manual_hello))
         )
         .service(
            web::scope("/app")
            .guard(guard::Host("www.rust-lang.org"))
            .route("", web::to(|| async { HttpResponse::Ok().body("www") })),
         )

    }).bind(("127.0.0.1", 7878))?
    .run()
    .await
}

