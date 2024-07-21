from datetime import datetime
from sqlalchemy import create_engine
from sqlalchemy.orm import scoped_session, sessionmaker
from sqlalchemy.orm.exc import NoResultFound
from goo_gl_archives.model.base import Base
from goo_gl_archives.model.link import Link
from typing import List, Dict


def init_sqlalchemy(database: str) -> scoped_session:
    """
    Initialize SQLAlchemy engine, create tables, and return a scoped session.

    Args:
        database (str): The database URL.

    Returns:
        scoped_session: The SQLAlchemy session object.
    """
    # Create DB Engine
    engine = create_engine(database, echo=False)

    # Create Tables
    Base.metadata.create_all(bind=engine)

    # Create DB Session
    session = scoped_session(
        sessionmaker(autocommit=False, autoflush=False, bind=engine)
    )

    return session


def insert_data(session: scoped_session, results: List[Dict[str, str]]) -> None:
    """
    Insert or update records in the database based on the given results.

    Args:
        session (scoped_session): The SQLAlchemy session object.
        results (list): A list of dictionaries containing link data.

    Returns:
        None
    """
    for result in results:
        try:
            # Search for a record that matches the uid
            link = session.query(Link).filter_by(uid=result["uid"]).one()

            # Skip if the record exists and all fields match
            if (
                link.redirect_url == result["redirect_url"]
                and link.domain_name == result["domain_name"]
                and link.site_title == result["site_title"]
                and link.http_status == result["http_status"]
            ):
                continue

            # Update if any field has changed
            link.redirect_url == result["redirect_url"]
            link.domain_name == result["domain_name"]
            link.site_title == result["site_title"]
            link.http_status == result["http_status"]
            link.updated_at = datetime.now()
        except NoResultFound:
            # Create a new record if it does not exist
            link = Link(
                uid=result["uid"],
                redirect_url=result["redirect_url"],
                domain_name=result["domain_name"],
                site_title=result["site_title"],
                http_status=result["http_status"],
            )
            session.add(link)

    session.commit()
