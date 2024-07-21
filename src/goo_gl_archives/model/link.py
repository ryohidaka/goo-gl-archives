from datetime import datetime
from sqlalchemy import Column, Integer, DateTime, String

from goo_gl_archives.model.base import Base


class Link(Base):
    """
    Link Model
    """

    __tablename__ = "links"

    id = Column(Integer, primary_key=True, autoincrement=True, unique=True)
    uid = Column(String(10), unique=True, nullable=False)
    redirect_url = Column(String)
    domain_name = Column(String)
    site_title = Column(String)
    http_status = Column(Integer)
    created_at = Column(DateTime, default=datetime.now, nullable=False)
    updated_at = Column(DateTime, default=datetime.now, nullable=False)
