import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

interface IAuthor {
  ID: number,
  firstName: string,
  middleName: string,
  lastName: string,
}

interface IBook {
  ID: number,
  authors: IAuthor[],
  title: string,
  copyrightYear: number,
  about: string,
  imgUrl: string,
}

@Component({
  selector: 'app-book',
  templateUrl: './book.component.html',
  styleUrls: ['./book.component.css']
})

export class BookComponent implements OnInit {
  public authors: IAuthor[] = [];
  public books: IBook[] = [];
  public author: any;

  constructor(private httpClient: HttpClient) { }

  ngOnInit() {
    this.getBooks();
  }

  getBooks() {
    this.httpClient.get<IBook[]>('/api/books')
    .subscribe({
      next: data => { 
        this.books = data;
      },
      error: error => {
        console.log('There was an error!', error);
      }
    })
  }

  // getAuthors() {
  //   this.httpClient.get<IAuthor[]>('/api/authors')
  //   .subscribe({
  //     next: data => { 
  //       this.authors = data;
  //     },
  //     error: error => {
  //       console.log('There was an error!', error);
  //     }
  //   })
  // }

  // getAuthorById(id: number) {
  //   this.httpClient.get<IAuthor[]>(`/api/authors/${id}`)
  //   .subscribe({
  //     next: data => { 
  //       this.author = data;
  //     },
  //     error: error => {
  //       console.log('There was an error!', error);
  //     }
  //   })
  // }

}
