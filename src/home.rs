use actix_web::{get, web, HttpResponse};
use handlebars::Handlebars;
use pulldown_cmark::{html, Options, Parser};
use std::fs;

// Path: /, Method: GE
#[get("/")]
pub async fn handle(hb: web::Data<Handlebars<'_>>) -> HttpResponse {
    // Create the markdown parser's options.
    let mut options = Options::empty();
    options.insert(Options::ENABLE_STRIKETHROUGH);

    // Read the sections dir.
    let dir = fs::read_dir("sections").unwrap();
    let filenames = dir
        .map(|entry| {
            // Get the filenames.
            // TODO: Improve this.
            format!(
                "sections/{}",
                String::from(entry.unwrap().path().file_name().unwrap().to_str().unwrap())
            )
        })
        .collect::<Vec<String>>();

    let mut html_output = vec![];
    for filename in filenames {
        // Let's parse the markdown.
        let content = fs::read_to_string(filename).expect("File in sections dir couldn't be read.");
        let parser = Parser::new_ext(&*content, options);

        // Push to the output vector.
        let mut parsed_content = String::new();
        html::push_html(&mut parsed_content, parser);
        html_output.push(parsed_content);
    }

    // Render the page with handlebars.
    let data = json!({
        "title": "The Fifth Site",
        "sections": html_output,
    });
    let body = hb.render("index", &data).unwrap();

    // 200 OK
    HttpResponse::Ok().body(body)
}
